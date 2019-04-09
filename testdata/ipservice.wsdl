<?xml version="1.0" encoding="utf-8"?>
<wsdl:definitions xmlns:s="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://schemas.xmlsoap.org/wsdl/soap12/" xmlns:http="http://schemas.xmlsoap.org/wsdl/http/" xmlns:mime="http://schemas.xmlsoap.org/wsdl/mime/" xmlns:tns="http://lavasoft.com/" xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/" xmlns:tm="http://microsoft.com/wsdl/mime/textMatching/" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/" targetNamespace="http://lavasoft.com/" xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">
  <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;b&gt;A web service which performs GetIpAddress Lookups.&lt;/b&gt;</wsdl:documentation>
  <wsdl:types>
    <s:schema elementFormDefault="qualified" targetNamespace="http://lavasoft.com/">
      <s:element name="GetIpLocation">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="sIp" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="GetIpLocationResponse">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="GetIpLocationResult" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="GetIpLocation_2_0">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="sIp" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="GetIpLocation_2_0Response">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="GetIpLocation_2_0Result" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="GetLocation">
        <s:complexType />
      </s:element>
      <s:element name="GetLocationResponse">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="GetLocationResult" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="GetCountryISO2ByName">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="countryName" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="GetCountryISO2ByNameResponse">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="GetCountryISO2ByNameResult" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="GetCountryNameByISO2">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="iso2Code" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="GetCountryNameByISO2Response">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="GetCountryNameByISO2Result" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="string" nillable="true" type="s:string" />
    </s:schema>
  </wsdl:types>
  <wsdl:message name="GetIpLocationSoapIn">
    <wsdl:part name="parameters" element="tns:GetIpLocation" />
  </wsdl:message>
  <wsdl:message name="GetIpLocationSoapOut">
    <wsdl:part name="parameters" element="tns:GetIpLocationResponse" />
  </wsdl:message>
  <wsdl:message name="GetIpLocation_2_0SoapIn">
    <wsdl:part name="parameters" element="tns:GetIpLocation_2_0" />
  </wsdl:message>
  <wsdl:message name="GetIpLocation_2_0SoapOut">
    <wsdl:part name="parameters" element="tns:GetIpLocation_2_0Response" />
  </wsdl:message>
  <wsdl:message name="GetLocationSoapIn">
    <wsdl:part name="parameters" element="tns:GetLocation" />
  </wsdl:message>
  <wsdl:message name="GetLocationSoapOut">
    <wsdl:part name="parameters" element="tns:GetLocationResponse" />
  </wsdl:message>
  <wsdl:message name="GetCountryISO2ByNameSoapIn">
    <wsdl:part name="parameters" element="tns:GetCountryISO2ByName" />
  </wsdl:message>
  <wsdl:message name="GetCountryISO2ByNameSoapOut">
    <wsdl:part name="parameters" element="tns:GetCountryISO2ByNameResponse" />
  </wsdl:message>
  <wsdl:message name="GetCountryNameByISO2SoapIn">
    <wsdl:part name="parameters" element="tns:GetCountryNameByISO2" />
  </wsdl:message>
  <wsdl:message name="GetCountryNameByISO2SoapOut">
    <wsdl:part name="parameters" element="tns:GetCountryNameByISO2Response" />
  </wsdl:message>
  <wsdl:message name="GetIpLocationHttpGetIn">
    <wsdl:part name="sIp" type="s:string" />
  </wsdl:message>
  <wsdl:message name="GetIpLocationHttpGetOut">
    <wsdl:part name="Body" element="tns:string" />
  </wsdl:message>
  <wsdl:message name="GetIpLocation_2_0HttpGetIn">
    <wsdl:part name="sIp" type="s:string" />
  </wsdl:message>
  <wsdl:message name="GetIpLocation_2_0HttpGetOut">
    <wsdl:part name="Body" element="tns:string" />
  </wsdl:message>
  <wsdl:message name="GetLocationHttpGetIn" />
  <wsdl:message name="GetLocationHttpGetOut">
    <wsdl:part name="Body" element="tns:string" />
  </wsdl:message>
  <wsdl:message name="GetCountryISO2ByNameHttpGetIn">
    <wsdl:part name="countryName" type="s:string" />
  </wsdl:message>
  <wsdl:message name="GetCountryISO2ByNameHttpGetOut">
    <wsdl:part name="Body" element="tns:string" />
  </wsdl:message>
  <wsdl:message name="GetCountryNameByISO2HttpGetIn">
    <wsdl:part name="iso2Code" type="s:string" />
  </wsdl:message>
  <wsdl:message name="GetCountryNameByISO2HttpGetOut">
    <wsdl:part name="Body" element="tns:string" />
  </wsdl:message>
  <wsdl:message name="GetIpLocationHttpPostIn">
    <wsdl:part name="sIp" type="s:string" />
  </wsdl:message>
  <wsdl:message name="GetIpLocationHttpPostOut">
    <wsdl:part name="Body" element="tns:string" />
  </wsdl:message>
  <wsdl:message name="GetIpLocation_2_0HttpPostIn">
    <wsdl:part name="sIp" type="s:string" />
  </wsdl:message>
  <wsdl:message name="GetIpLocation_2_0HttpPostOut">
    <wsdl:part name="Body" element="tns:string" />
  </wsdl:message>
  <wsdl:message name="GetLocationHttpPostIn" />
  <wsdl:message name="GetLocationHttpPostOut">
    <wsdl:part name="Body" element="tns:string" />
  </wsdl:message>
  <wsdl:message name="GetCountryISO2ByNameHttpPostIn">
    <wsdl:part name="countryName" type="s:string" />
  </wsdl:message>
  <wsdl:message name="GetCountryISO2ByNameHttpPostOut">
    <wsdl:part name="Body" element="tns:string" />
  </wsdl:message>
  <wsdl:message name="GetCountryNameByISO2HttpPostIn">
    <wsdl:part name="iso2Code" type="s:string" />
  </wsdl:message>
  <wsdl:message name="GetCountryNameByISO2HttpPostOut">
    <wsdl:part name="Body" element="tns:string" />
  </wsdl:message>
  <wsdl:portType name="GeoIPServiceSoap">
    <wsdl:operation name="GetIpLocation">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country and State XML String for a given ip address in x.x.x.x format&lt;br/&gt;&lt;br/&gt;Parameter:GetIpAddress address &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;Return:format in XML &lt; GeoIP &gt;&lt; Country &gt;US&lt; /Country &gt;&lt; State &gt;PA&lt; /State &gt;&lt; /GeoIP &gt; &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;</wsdl:documentation>
      <wsdl:input message="tns:GetIpLocationSoapIn" />
      <wsdl:output message="tns:GetIpLocationSoapOut" />
    </wsdl:operation>
    <wsdl:operation name="GetIpLocation_2_0">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country and State XML String for a given ip address in x.x.x.x format&lt;br/&gt;&lt;br/&gt;Parameter:GetIpAddress address &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;Return:format in XML &lt; GeoIP &gt;&lt; Country &gt;US&lt; /Country &gt;&lt; State &gt;PA&lt; /State &gt;&lt; /GeoIP &gt; &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt; or &lt;b&gt;string empty&lt;/b&gt; </wsdl:documentation>
      <wsdl:input message="tns:GetIpLocation_2_0SoapIn" />
      <wsdl:output message="tns:GetIpLocation_2_0SoapOut" />
    </wsdl:operation>
    <wsdl:operation name="GetLocation">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country and State XML String for a given ip address in x.x.x.x format&lt;br/&gt;&lt;br/&gt;Parameter:GetIpAddress address &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;Return:format in XML &lt; GeoIP &gt;&lt; Country &gt;US&lt; /Country &gt;&lt; State &gt;PA&lt; /State &gt;&lt; /GeoIP &gt; &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;</wsdl:documentation>
      <wsdl:input message="tns:GetLocationSoapIn" />
      <wsdl:output message="tns:GetLocationSoapOut" />
    </wsdl:operation>
    <wsdl:operation name="GetCountryISO2ByName">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country ISO2 code by Country Name&lt;br/&gt;Paramater:Country Name &amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;b/&gt;</wsdl:documentation>
      <wsdl:input message="tns:GetCountryISO2ByNameSoapIn" />
      <wsdl:output message="tns:GetCountryISO2ByNameSoapOut" />
    </wsdl:operation>
    <wsdl:operation name="GetCountryNameByISO2">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country Name by Country ISO2 code&lt;br/&gt;Paramater:Country Code &amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;b/&gt;&lt;br/&gt;Country code  format is ALPHA-2 (CA - for Canada)</wsdl:documentation>
      <wsdl:input message="tns:GetCountryNameByISO2SoapIn" />
      <wsdl:output message="tns:GetCountryNameByISO2SoapOut" />
    </wsdl:operation>
  </wsdl:portType>
  <wsdl:portType name="GeoIPServiceHttpGet">
    <wsdl:operation name="GetIpLocation">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country and State XML String for a given ip address in x.x.x.x format&lt;br/&gt;&lt;br/&gt;Parameter:GetIpAddress address &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;Return:format in XML &lt; GeoIP &gt;&lt; Country &gt;US&lt; /Country &gt;&lt; State &gt;PA&lt; /State &gt;&lt; /GeoIP &gt; &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;</wsdl:documentation>
      <wsdl:input message="tns:GetIpLocationHttpGetIn" />
      <wsdl:output message="tns:GetIpLocationHttpGetOut" />
    </wsdl:operation>
    <wsdl:operation name="GetIpLocation_2_0">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country and State XML String for a given ip address in x.x.x.x format&lt;br/&gt;&lt;br/&gt;Parameter:GetIpAddress address &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;Return:format in XML &lt; GeoIP &gt;&lt; Country &gt;US&lt; /Country &gt;&lt; State &gt;PA&lt; /State &gt;&lt; /GeoIP &gt; &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt; or &lt;b&gt;string empty&lt;/b&gt; </wsdl:documentation>
      <wsdl:input message="tns:GetIpLocation_2_0HttpGetIn" />
      <wsdl:output message="tns:GetIpLocation_2_0HttpGetOut" />
    </wsdl:operation>
    <wsdl:operation name="GetLocation">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country and State XML String for a given ip address in x.x.x.x format&lt;br/&gt;&lt;br/&gt;Parameter:GetIpAddress address &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;Return:format in XML &lt; GeoIP &gt;&lt; Country &gt;US&lt; /Country &gt;&lt; State &gt;PA&lt; /State &gt;&lt; /GeoIP &gt; &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;</wsdl:documentation>
      <wsdl:input message="tns:GetLocationHttpGetIn" />
      <wsdl:output message="tns:GetLocationHttpGetOut" />
    </wsdl:operation>
    <wsdl:operation name="GetCountryISO2ByName">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country ISO2 code by Country Name&lt;br/&gt;Paramater:Country Name &amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;b/&gt;</wsdl:documentation>
      <wsdl:input message="tns:GetCountryISO2ByNameHttpGetIn" />
      <wsdl:output message="tns:GetCountryISO2ByNameHttpGetOut" />
    </wsdl:operation>
    <wsdl:operation name="GetCountryNameByISO2">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country Name by Country ISO2 code&lt;br/&gt;Paramater:Country Code &amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;b/&gt;&lt;br/&gt;Country code  format is ALPHA-2 (CA - for Canada)</wsdl:documentation>
      <wsdl:input message="tns:GetCountryNameByISO2HttpGetIn" />
      <wsdl:output message="tns:GetCountryNameByISO2HttpGetOut" />
    </wsdl:operation>
  </wsdl:portType>
  <wsdl:portType name="GeoIPServiceHttpPost">
    <wsdl:operation name="GetIpLocation">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country and State XML String for a given ip address in x.x.x.x format&lt;br/&gt;&lt;br/&gt;Parameter:GetIpAddress address &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;Return:format in XML &lt; GeoIP &gt;&lt; Country &gt;US&lt; /Country &gt;&lt; State &gt;PA&lt; /State &gt;&lt; /GeoIP &gt; &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;</wsdl:documentation>
      <wsdl:input message="tns:GetIpLocationHttpPostIn" />
      <wsdl:output message="tns:GetIpLocationHttpPostOut" />
    </wsdl:operation>
    <wsdl:operation name="GetIpLocation_2_0">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country and State XML String for a given ip address in x.x.x.x format&lt;br/&gt;&lt;br/&gt;Parameter:GetIpAddress address &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;Return:format in XML &lt; GeoIP &gt;&lt; Country &gt;US&lt; /Country &gt;&lt; State &gt;PA&lt; /State &gt;&lt; /GeoIP &gt; &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt; or &lt;b&gt;string empty&lt;/b&gt; </wsdl:documentation>
      <wsdl:input message="tns:GetIpLocation_2_0HttpPostIn" />
      <wsdl:output message="tns:GetIpLocation_2_0HttpPostOut" />
    </wsdl:operation>
    <wsdl:operation name="GetLocation">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country and State XML String for a given ip address in x.x.x.x format&lt;br/&gt;&lt;br/&gt;Parameter:GetIpAddress address &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;Return:format in XML &lt; GeoIP &gt;&lt; Country &gt;US&lt; /Country &gt;&lt; State &gt;PA&lt; /State &gt;&lt; /GeoIP &gt; &amp;nbsp;&amp;nbsp;&amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;/b&gt;</wsdl:documentation>
      <wsdl:input message="tns:GetLocationHttpPostIn" />
      <wsdl:output message="tns:GetLocationHttpPostOut" />
    </wsdl:operation>
    <wsdl:operation name="GetCountryISO2ByName">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country ISO2 code by Country Name&lt;br/&gt;Paramater:Country Name &amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;b/&gt;</wsdl:documentation>
      <wsdl:input message="tns:GetCountryISO2ByNameHttpPostIn" />
      <wsdl:output message="tns:GetCountryISO2ByNameHttpPostOut" />
    </wsdl:operation>
    <wsdl:operation name="GetCountryNameByISO2">
      <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;br/&gt;Get Country Name by Country ISO2 code&lt;br/&gt;Paramater:Country Code &amp;nbsp;&amp;nbsp;&lt;b&gt;Type:String&lt;b/&gt;&lt;br/&gt;Country code  format is ALPHA-2 (CA - for Canada)</wsdl:documentation>
      <wsdl:input message="tns:GetCountryNameByISO2HttpPostIn" />
      <wsdl:output message="tns:GetCountryNameByISO2HttpPostOut" />
    </wsdl:operation>
  </wsdl:portType>
  <wsdl:binding name="GeoIPServiceSoap" type="tns:GeoIPServiceSoap">
    <soap:binding transport="http://schemas.xmlsoap.org/soap/http" />
    <wsdl:operation name="GetIpLocation">
      <soap:operation soapAction="http://lavasoft.com/GetIpLocation" style="document" />
      <wsdl:input>
        <soap:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetIpLocation_2_0">
      <soap:operation soapAction="http://lavasoft.com/GetIpLocation_2_0" style="document" />
      <wsdl:input>
        <soap:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetLocation">
      <soap:operation soapAction="http://lavasoft.com/GetLocation" style="document" />
      <wsdl:input>
        <soap:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetCountryISO2ByName">
      <soap:operation soapAction="http://lavasoft.com/GetCountryISO2ByName" style="document" />
      <wsdl:input>
        <soap:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetCountryNameByISO2">
      <soap:operation soapAction="http://lavasoft.com/GetCountryNameByISO2" style="document" />
      <wsdl:input>
        <soap:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
  </wsdl:binding>
  <wsdl:binding name="GeoIPServiceSoap12" type="tns:GeoIPServiceSoap">
    <soap12:binding transport="http://schemas.xmlsoap.org/soap/http" />
    <wsdl:operation name="GetIpLocation">
      <soap12:operation soapAction="http://lavasoft.com/GetIpLocation" style="document" />
      <wsdl:input>
        <soap12:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap12:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetIpLocation_2_0">
      <soap12:operation soapAction="http://lavasoft.com/GetIpLocation_2_0" style="document" />
      <wsdl:input>
        <soap12:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap12:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetLocation">
      <soap12:operation soapAction="http://lavasoft.com/GetLocation" style="document" />
      <wsdl:input>
        <soap12:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap12:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetCountryISO2ByName">
      <soap12:operation soapAction="http://lavasoft.com/GetCountryISO2ByName" style="document" />
      <wsdl:input>
        <soap12:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap12:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetCountryNameByISO2">
      <soap12:operation soapAction="http://lavasoft.com/GetCountryNameByISO2" style="document" />
      <wsdl:input>
        <soap12:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap12:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
  </wsdl:binding>
  <wsdl:binding name="GeoIPServiceHttpGet" type="tns:GeoIPServiceHttpGet">
    <http:binding verb="GET" />
    <wsdl:operation name="GetIpLocation">
      <http:operation location="/GetIpLocation" />
      <wsdl:input>
        <http:urlEncoded />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetIpLocation_2_0">
      <http:operation location="/GetIpLocation_2_0" />
      <wsdl:input>
        <http:urlEncoded />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetLocation">
      <http:operation location="/GetLocation" />
      <wsdl:input>
        <http:urlEncoded />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetCountryISO2ByName">
      <http:operation location="/GetCountryISO2ByName" />
      <wsdl:input>
        <http:urlEncoded />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetCountryNameByISO2">
      <http:operation location="/GetCountryNameByISO2" />
      <wsdl:input>
        <http:urlEncoded />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
  </wsdl:binding>
  <wsdl:binding name="GeoIPServiceHttpPost" type="tns:GeoIPServiceHttpPost">
    <http:binding verb="POST" />
    <wsdl:operation name="GetIpLocation">
      <http:operation location="/GetIpLocation" />
      <wsdl:input>
        <mime:content type="application/x-www-form-urlencoded" />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetIpLocation_2_0">
      <http:operation location="/GetIpLocation_2_0" />
      <wsdl:input>
        <mime:content type="application/x-www-form-urlencoded" />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetLocation">
      <http:operation location="/GetLocation" />
      <wsdl:input>
        <mime:content type="application/x-www-form-urlencoded" />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetCountryISO2ByName">
      <http:operation location="/GetCountryISO2ByName" />
      <wsdl:input>
        <mime:content type="application/x-www-form-urlencoded" />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="GetCountryNameByISO2">
      <http:operation location="/GetCountryNameByISO2" />
      <wsdl:input>
        <mime:content type="application/x-www-form-urlencoded" />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
  </wsdl:binding>
  <wsdl:service name="GeoIPService">
    <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">&lt;b&gt;A web service which performs GetIpAddress Lookups.&lt;/b&gt;</wsdl:documentation>
    <wsdl:port name="GeoIPServiceSoap" binding="tns:GeoIPServiceSoap">
      <soap:address location="http://wsgeoip.lavasoft.com/ipservice.asmx" />
    </wsdl:port>
    <wsdl:port name="GeoIPServiceSoap12" binding="tns:GeoIPServiceSoap12">
      <soap12:address location="http://wsgeoip.lavasoft.com/ipservice.asmx" />
    </wsdl:port>
    <wsdl:port name="GeoIPServiceHttpGet" binding="tns:GeoIPServiceHttpGet">
      <http:address location="http://wsgeoip.lavasoft.com/ipservice.asmx" />
    </wsdl:port>
    <wsdl:port name="GeoIPServiceHttpPost" binding="tns:GeoIPServiceHttpPost">
      <http:address location="http://wsgeoip.lavasoft.com/ipservice.asmx" />
    </wsdl:port>
  </wsdl:service>
</wsdl:definitions>